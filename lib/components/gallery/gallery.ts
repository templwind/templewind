import './gallery.scss';

export class TwGallery extends HTMLElement {
	constructor() {
		super();
	}

	connectedCallback(): void {
		console.log("Gallery connected");
	}
}

customElements.define("tw-gallery", TwGallery);