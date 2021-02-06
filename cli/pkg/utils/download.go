package utils

import (
	"context"
	"github.com/mikamai/karavel/cli/pkg/logger"
	"github.com/pkg/errors"
	"io"
	"net/http"
	"os"
	"path"
	"time"
)

// Thank you Cosmin!
// https://gist.github.com/albulescu/e61979cc852e4ee8f49c

func DownloadWithProgress(ctx context.Context, log logger.Logger, url string, filename string) error {
	file := path.Base(url)

	log.Infof("Downloading file %s from %s\n", file, url)
	log.Info()

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
	//go printDownloadPercent(ctx, log, done, filename, int64(size))

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
	log.Info()
	log.Infof("Download completed in %s", elapsed)
	return nil
}

func printDownloadPercent(ctx context.Context, log logger.Logger, done chan int64, path string, total int64) {
	for {
		select {
		case <-ctx.Done():
			log.Error(ctx.Err())
			return
		case <-done:
			return
		default:
			file, err := os.Open(path)
			if err != nil {
				log.Error(err)
				return
			}

			fi, err := file.Stat()
			if err != nil {
				log.Error(err)
				return
			}

			size := fi.Size()
			if size == 0 {
				size = 1
			}

			percent := float64(size) / float64(total) * 100

			log.Infof("%.0f%%", percent)
		}
	}
}
