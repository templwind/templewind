import './file-input.scss';

export class TwFileInput extends HTMLElement {
	constructor() {
		super();
	}

	connectedCallback(): void {
		console.log("FileInput connected");
	}
}

customElements.define("tw-file-input", TwFileInput);