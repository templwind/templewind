import './input-field.scss';

export class TwInputField extends HTMLElement {
	constructor() {
		super();
	}

	connectedCallback(): void {
		console.log("InputField connected");
	}
}

customElements.define("tw-input-field", TwInputField);