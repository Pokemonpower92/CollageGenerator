package datastore

import (
	"fmt"
	"image"
	"image/draw"
	_ "image/png"
	"io"
	"log"
	"os"
	"path/filepath"

	"github.com/google/uuid"
)

type LocalStore struct {
	Directory string
	logger    *log.Logger
}

func NewLocalStore(directory string) *LocalStore {
	return &LocalStore{
		Directory: directory,
		logger:    log.New(log.Writer(), "", log.LstdFlags),
	}
}

func imageToRGBA(src image.Image) *image.RGBA {
	if dst, ok := src.(*image.RGBA); ok {
		return dst
	}
	b := src.Bounds()
	dst := image.NewRGBA(image.Rect(0, 0, b.Dx(), b.Dy()))
	draw.Draw(dst, dst.Bounds(), src, b.Min, draw.Src)
	return dst
}

func (s *LocalStore) GetImage(id uuid.UUID) (*image.RGBA, error) {
	fileName := fmt.Sprintf("%s.png", id.String())
	s.logger.Printf("Reading images from directory: %s", s.Directory)
	path := fmt.Sprintf("%s/%s", s.Directory, fileName)
	s.logger.Printf("Decoding file %s", path)
	f, err := os.Open(path)
	if err != nil {
		s.logger.Printf("Failed to open file: %s", fileName)
		return nil, err
	}
	defer f.Close()
	im, _, err := image.Decode(f)
	if err != nil {
		s.logger.Printf("Failed to decode image: %s", err)
		return nil, err
	}
	b := im.Bounds()
	rgba := image.NewRGBA(image.Rect(0, 0, b.Dx(), b.Dy()))
	draw.Draw(rgba, rgba.Bounds(), im, b.Min, draw.Src)
	return rgba, nil
}

func (s *LocalStore) PutImage(id uuid.UUID, reader io.Reader) error {
	dst, err := os.Create(filepath.Join(s.Directory, id.String()))
	if err != nil {
		return err
	}
	defer dst.Close()
	s.logger.Printf("Created image destination: %s", dst.Name())
	if _, err := io.Copy(dst, reader); err != nil {
		return err
	}
	s.logger.Printf("Successfully stored image.")
	return nil
}
