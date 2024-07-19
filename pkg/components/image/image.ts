import './image.scss';

export class TwImage extends HTMLElement {
	constructor() {
		super();
	}

	connectedCallback(): void {
		console.log("Image connected");
	}
}

customElements.define("tw-image", TwImage);