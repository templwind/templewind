import './paragraph.scss';

export class TwParagraph extends HTMLElement {
	constructor() {
		super();
	}

	connectedCallback(): void {
		console.log("Paragraph connected");
	}
}

customElements.define("tw-paragraph", TwParagraph);