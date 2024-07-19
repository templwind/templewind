import './rating.scss';

export class TwRating extends HTMLElement {
	constructor() {
		super();
	}

	connectedCallback(): void {
		console.log("Rating connected");
	}
}

customElements.define("tw-rating", TwRating);