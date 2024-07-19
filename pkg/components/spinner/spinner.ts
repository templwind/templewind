import './spinner.scss';

export class TwSpinner extends HTMLElement {
	constructor() {
		super();
	}

	connectedCallback(): void {
		console.log("Spinner connected");
	}
}

customElements.define("tw-spinner", TwSpinner);