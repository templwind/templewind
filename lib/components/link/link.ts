import './link.scss';

export class TwLink extends HTMLElement {
	constructor() {
		super();
	}

	connectedCallback(): void {
		console.log("Link connected");
	}
}

customElements.define("tw-link", TwLink);