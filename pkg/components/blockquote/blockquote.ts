import './blockquote.scss';

export class TwBlockquote extends HTMLElement {
	constructor() {
		super();
	}

	connectedCallback(): void {
		console.log("Blockquote connected");
	}
}

customElements.define("tw-blockquote", TwBlockquote);