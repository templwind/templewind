import './card.scss';

export class TwCard extends HTMLElement {
	constructor() {
		super();
	}

	connectedCallback(): void {
		console.log("Card connected");
	}
}

customElements.define("tw-card", TwCard);