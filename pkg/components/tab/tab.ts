import './tab.scss';

export class TwTab extends HTMLElement {
	constructor() {
		super();
	}

	connectedCallback(): void {
		console.log("Tab connected");
	}
}

customElements.define("tw-tab", TwTab);