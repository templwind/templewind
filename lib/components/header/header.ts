import './header.scss';

export class TwHeader extends HTMLElement {
	constructor() {
		super();
	}

	connectedCallback(): void {
		console.log("Header connected");
	}
}

customElements.define("tw-header", TwHeader);