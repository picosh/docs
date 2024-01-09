document.addEventListener("DOMContentLoaded", init);

function topNav() {
  const topLevel = document.querySelectorAll(".nav-list li");
  topLevel.forEach((li) => {
    const link = li.querySelector("a");
    const href = link.getAttribute("href");
    if (location.pathname === href) {
      li.classList.add("current-page");
    }
  });
}

function init() {
  topNav();

  let throttle = null;
  const headers = document.querySelectorAll("h2,h3,h4");
  const scrollEl = document.querySelector(".content");
  scrollEl.addEventListener("scroll", updateNav, { passive: true });
  window.addEventListener("load", updateNav);

  const nav = document.querySelector(".toc");
  const btns = document.querySelectorAll(".toc-btn");
  btns.forEach((btn) => {
    btn.addEventListener("click", (e) => {
      e.preventDefault();
      nav.classList.toggle("nav-show");
    });
  });

  function updateNav() {
    let h2, h3, h4;
    for (let i = 0; i < headers.length; i++) {
      const h = headers[i];
      const top = h.getBoundingClientRect().top;
      if (top > 10) {
        break;
      }

      if (h.tagName === "H2") {
        h2 = h;
        h3 = h4 = null;
      } else if (h.tagName === "H3") {
        h3 = h;
        h4 = null;
      } else {
        h4 = h;
      }
    }

    for (let i = 0; i < headers.length; i++) {
      const h = headers[i];
      if (h.tagName === "H4") {
        continue;
      }
      const nav = document.getElementById("nav-" + h.id);
      if (nav) {
        nav.classList.toggle("current", h === (h3 || h2));
      }
    }

    // Throttle to avoid crashes in Safari
    const h = h4 || h3 || h2;
    pathhash = location.pathname + (h ? "#" + h.id : "");
    if (throttle === null) {
      throttle = setTimeout(updatePathname, 300);
    }
  }

  function updatePathname() {
    throttle = null;
    if (location.pathname + location.hash !== pathhash) {
      history.replaceState(null, "", pathhash);
    }
  }
}
