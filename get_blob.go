package main

func (s *server) GetBlob(sha256 string) ([]byte, error) {
	return s.storage.Read(sha256)
}
