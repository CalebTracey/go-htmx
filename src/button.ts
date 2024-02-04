const btn = document.querySelector("button");

function random(number: number) {
    return Math.floor(Math.random() * (number + 1));
}

export function aboutHandler(e: Event)  {
    console.log("ABOUT")
}

// @ts-ignore
btn.addEventListener("click", () => {
    document.body.style.backgroundColor = `rgb(${random(255)} ${random(255)} ${random(255)})`;
});