class FileInput extends HTMLElement {
	constructor() {
		super();
	}

	connectedCallback(): void {
		console.log("FileInput connected");
	}
}

customElements.define("tw-file-input", FileInput);