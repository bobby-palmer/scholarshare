package main

import "mime/multipart"

func validateFile(file multipart.File, header *multipart.FileHeader) bool {
  return true
}
