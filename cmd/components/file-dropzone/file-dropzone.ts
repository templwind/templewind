class FileDropzone extends HTMLElement {
	constructor() {
		super();
	}

	connectedCallback(): void {
		console.log("FileDropzone connected");
	}
}

customElements.define("tw-file-dropzone", FileDropzone);