import './checkbox.scss';

export class TwCheckbox extends HTMLElement {
	constructor() {
		super();
	}

	connectedCallback(): void {
		console.log("Checkbox connected");
	}
}

customElements.define("tw-checkbox", TwCheckbox);