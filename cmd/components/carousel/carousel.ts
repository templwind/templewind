class Carousel extends HTMLElement {
	constructor() {
		super();
	}

	connectedCallback(): void {
		console.log("Carousel connected");
	}
}

customElements.define("tw-carousel", Carousel);