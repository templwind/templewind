import './alert.scss';

export class TwAlert extends HTMLElement {
	constructor() {
		super();
	}

	connectedCallback(): void {
		console.log("Alert connected");
	}
}

customElements.define("tw-alert", TwAlert);