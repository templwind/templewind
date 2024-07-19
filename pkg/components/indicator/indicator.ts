import './indicator.scss';

export class TwIndicator extends HTMLElement {
	constructor() {
		super();
	}

	connectedCallback(): void {
		console.log("Indicator connected");
	}
}

customElements.define("tw-indicator", TwIndicator);