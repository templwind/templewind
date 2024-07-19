import './footer.scss';

export class TwFooter extends HTMLElement {
	constructor() {
		super();
	}

	connectedCallback(): void {
		console.log("Footer connected");
	}
}

customElements.define("tw-footer", TwFooter);