import './file-dropzone.scss';

export class TwFileDropzone extends HTMLElement {
	constructor() {
		super();
	}

	connectedCallback(): void {
		console.log("FileDropzone connected");
	}
}

customElements.define("tw-file-dropzone", TwFileDropzone);