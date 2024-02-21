package main

import (
	"bytes"
	"database/sql"
	"image"
	"log"
	"math/rand"
	"net/http"
	"os"
	"path"
	"time"

	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"

	"github.com/chai2010/webp"

	_ "github.com/mattn/go-sqlite3"

	"keiko/keikodb"
)

type WebpImage = bytes.Buffer

// Scans the given directory and returns all entries.
func getImageNames(path string) []os.DirEntry {
	entries, err := os.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}
	return entries
}

// Read raw image data from disk.
func loadImg(path string) []byte {
	data, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	return data
}

// Parse the raw bytes of an image into an Image format.
func parseImg(rawImg []byte) image.Image {
	img, _, err := image.Decode(bytes.NewReader(rawImg))
	if err != nil {
		log.Fatal(err)
	}
	return img
}

// Convert an image to WebP format.
func imgToWebp(img image.Image) WebpImage {
	var buf bytes.Buffer
	if err := webp.Encode(&buf, img, &webp.Options{Lossless: true}); err != nil {
		log.Fatal(err)
	}
	return buf
}

// Returns a random image from the `data` directory.
// The first return value is the name of the image,
// the second return value is the image data in WebP format.
func getImage() (string, WebpImage) {
	imgBase := "data"

	// find all the names of the images
	imgName := getImageNames(imgBase)
	// pick an index at random
	index := rand.Intn(len(imgName))

	// get the randomly picked image
	imgEntry := imgName[index]
	// construct a path to the image in the `data` directory
	imgPath := path.Join(imgBase, imgEntry.Name())
	// load the image into memory
	rawImg := loadImg(imgPath)

	// parse into an image.Image
	parsedImg := parseImg(rawImg)
	// convert to WebP
	webpImg := imgToWebp(parsedImg)

	return imgEntry.Name(), webpImg
}

// handler for a GET request on `/`
func handleRoot(db *sql.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		if req.Method != "GET" {
			return
		}

		// get a random image
		name, img := getImage()

		// increment the hit counter for this image
		keikodb.IncrementHitCount(db, name)

		// serve it
		log.Println("Serve", name)
		w.Write(img.Bytes())

		return
	})
}

func main() {
	rand.Seed(time.Now().UnixNano())

	log.Println("Using database 'keiko.db'")
	db, err := sql.Open("sqlite3", "keiko.db")
	if err != nil {
		log.Fatal(err)
	}

	// try to create a new database
	keikodb.MakeNew(db)

	// serve on `/`
	http.Handle("/", handleRoot(db))

	log.Println("Listening on localhost:8099")
	err = http.ListenAndServe("localhost:8099", nil)
	if err != nil {
		log.Fatalln(err)
	}
}
