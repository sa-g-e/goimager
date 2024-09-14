[![another banger typewriter effect](https://readme-typing-svg.demolab.com?font=Fira+Code&pause=1000&repeat=false&width=600&height=60&lines=goimager+-+a+minimal+image+uploader+written+in+go)](https://git.io/typing-svg)

**goimager** is an easy-to-setup, lightweight image uploader.

(My 2nd project in Go.) If you want to see something similar, check out [gobin](https://github.com/sa-g-e/gobin), a minimal paste bin written in Go.

## Features

- Easy image uploads with secret key authentication
- Supports common image formats (JPG, JPEG, PNG, GIF, BMP)
- Generates unique URLs for uploaded images
- Easy integration with ShareX

## Prerequisites

- Go 1.16 or higher
- ShareX (for screenshot uploading)

## Installation

1. **Clone the Repo**

Open your terminal and run:

```bash
git clone https://github.com/sa-g-e/goimager.git
cd goimager
```
   
2. **Set Your Secret Key**

Open `handlers.go` and go to the line:

`const secretKey = "YOUR_SECRET_KEY_HERE"`

Replace "YOUR_SECRET_KEY_HERE" with your actual secret key.

3. Run the Server

To start the server, run:

`go run main.go handlers.go utils.go`

4. Build the Server (Optional)

To build the server into an executable, use:

`go build`

Then run the executable:
`./goimager`

On Windows, run:

`goimager.exe`

By default, it'll run on `http://localhost:8080`

## ShareX Configuration
To upload screenshots to goimager using [ShareX](https://getsharex.com/), follow these steps:

Open ShareX and go to Destinations > Custom uploader settings.

Click New and then Import from clipboard.

Copy the following configuration and press OK:

```json
{
  "Version": "16.1.0",
  "Name": "goimager",
  "DestinationType": "ImageUploader",
  "RequestMethod": "POST",
  "RequestURL": "http://localhost:8080/upload",
  "Headers": {
    "X-Secret-Key": "test"
  },
  "Body": "MultipartFormData",
  "FileFormName": "file",
  "URL": "{json:url}"
}
```
Set your actual secret key in the X-Secret-Key field to match the one in handlers.go.
