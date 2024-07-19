import './floating-label.scss';

export class TwFloatingLabel extends HTMLElement {
	constructor() {
		super();
	}

	connectedCallback(): void {
		console.log("FloatingLabel connected");
	}
}

customElements.define("tw-floating-label", TwFloatingLabel);