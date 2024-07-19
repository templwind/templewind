import './button.scss';

export class TwButton extends HTMLElement {
	constructor() {
		super();
	}

	connectedCallback(): void {
		console.log("Button connected");
	}
}

customElements.define("tw-button", TwButton);