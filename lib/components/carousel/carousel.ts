import './carousel.scss';

export class TwCarousel extends HTMLElement {
	constructor() {
		super();
	}

	connectedCallback(): void {
		console.log("Carousel connected");
	}
}

customElements.define("tw-carousel", TwCarousel);