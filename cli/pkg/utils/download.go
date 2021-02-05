package utils

import (
	"context"
	"github.com/pkg/errors"
	"io"
	"log"
	"net/http"
	"os"
	"path"
	"time"
)

// Thank you Cosmin!
// https://gist.github.com/albulescu/e61979cc852e4ee8f49c

func DownloadWithProgress(ctx context.Context, logger *log.Logger, url string, filename string) error {
	file := path.Base(url)

	logger.Printf("Downloading file %s from %s\n", file, url)
	logger.Println()

	start := time.Now()

	out, err := os.OpenFile(filename, os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		return err
	}

	defer out.Close()

	headResp, err := http.Head(url)
	if err != nil {
		return err
	}

	defer headResp.Body.Close()
	if headResp.StatusCode >= 400 {
		return errors.Errorf("failed to fetch %s: %s", url, headResp.Status)
	}

	//size, err := strconv.Atoi(headResp.Header.Get("Content-Length"))
	//if err != nil {
	//    return err
	//}

	//done := make(chan int64)

	// TODO: fix printing
	//go printDownloadPercent(ctx, logger, done, filename, int64(size))

	resp, err := http.Get(url)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	//done <- n

	elapsed := time.Since(start)
	logger.Println()
	logger.Printf("Download completed in %s", elapsed)
	return nil
}

func printDownloadPercent(ctx context.Context, logger *log.Logger, done chan int64, path string, total int64) {
	for {
		select {
		case <-ctx.Done():
			logger.Println(ctx.Err())
			return
		case <-done:
			return
		default:
			file, err := os.Open(path)
			if err != nil {
				logger.Print(err)
				return
			}

			fi, err := file.Stat()
			if err != nil {
				logger.Print(err)
				return
			}

			size := fi.Size()
			if size == 0 {
				size = 1
			}

			percent := float64(size) / float64(total) * 100

			logger.Printf("%.0f%%", percent)
		}
	}
}
