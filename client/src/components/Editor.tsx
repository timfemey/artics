import { Component } from "solid-js";

const Editor: Component<{}> = () => {
  const blogTitleField = document.querySelector(".title");
  const articleFeild = document.querySelector(".article");

  // banner
  const bannerImage = document.querySelector("#banner-upload");
  const banner = document.querySelector(".banner");
  let bannerPath;

  const publishBtn = document.querySelector(".publish-btn");
  const uploadInput = document.querySelector("#image-upload");

  function uploadImage(uploadFile: any) {
    const [file] = uploadFile.files;
  }

  bannerImage?.addEventListener("change", () => {
    uploadImage(bannerImage);
  });

  uploadInput?.addEventListener("change", () => {
    uploadImage(uploadInput);
  });

  return (
    <>
      <div class="banner">
        <input type="file" accept="image/*" id="banner-upload" hidden />
        <label for="banner-upload" class="banner-upload-btn">
          <i class="fa-solid fa-upload"></i>
        </label>
      </div>
      <div class="blog">
        <textarea class="title" placeholder="Blog title..."></textarea>
        <textarea
          class="article"
          placeholder="Start writing here..."
        ></textarea>
      </div>
      <div class="blog-options">
        <button class="btn dark publish-btn">publish</button>
        <input type="file" accept="image/*" id="image-upload" hidden />
        <label for="image-upload" class="btn grey upload-btn">
          Upload Image
        </label>
      </div>
    </>
  );
};

export default Editor;
