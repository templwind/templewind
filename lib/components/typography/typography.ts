import './typography.scss';

export class TwTypography extends HTMLElement {
	constructor() {
		super();
	}

	connectedCallback(): void {
		console.log("Typography connected");
	}
}

customElements.define("tw-typography", TwTypography);