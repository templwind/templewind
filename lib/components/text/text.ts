import './text.scss';

export class TwText extends HTMLElement {
	constructor() {
		super();
	}

	connectedCallback(): void {
		console.log("Text connected");
	}
}

customElements.define("tw-text", TwText);