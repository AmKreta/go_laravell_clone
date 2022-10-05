package helper

import "os"

func CreateFile(path, name string) error {
	if _, err := os.Stat(path + "/" + name); os.IsNotExist(err) {
		file, err := os.Create(path + "/" + name)
		if err != nil {
			return err
		}
		defer func(file *os.File) {
			_ = file.Close()
		}(file)

	}
	return nil
}

func CreateFileWithHandler(path, name string, handler func(*os.File)) error {
	if _, err := os.Stat(path + "/" + name); os.IsNotExist(err) {
		file, err := os.Create(path + "/" + name)
		if err != nil {
			return err
		}
		defer func(file *os.File) {
			handler(file)
			_ = file.Close()
		}(file)

	}
	return nil
}
