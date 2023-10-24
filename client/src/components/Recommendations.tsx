import { Component } from "solid-js";

const Recommendations: Component<{}> = () => {
  return (
    <>
      <section class="blogs-section">
        <div class="blog-card">
          <img
            src="https://images.pexels.com/photos/842711/pexels-photo-842711.jpeg?auto=compress&cs=tinysrgb&w=1260&h=750&dpr=1"
            class="blog-image"
            alt="Article Banner"
          />
          <h1 class="blog-title">Lorem ipsum dolor sit amet consectetur.</h1>
          <p class="blog-overview">
            Lorem ipsum dolor sit amet consectetur adipisicing elit. Sunt
            incidunt fugiat quos porro repellat harum. Adipisci tempora corporis
            rem cum.
          </p>
          <a href="/" class="btn dark">
            read
          </a>
        </div>
      </section>
    </>
  );
};

export default Recommendations;
