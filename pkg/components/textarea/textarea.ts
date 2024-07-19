import './textarea.scss';

export class TwTextarea extends HTMLElement {
	constructor() {
		super();
	}

	connectedCallback(): void {
		console.log("Textarea connected");
	}
}

customElements.define("tw-textarea", TwTextarea);