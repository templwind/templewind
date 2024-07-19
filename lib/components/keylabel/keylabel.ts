import './keylabel.scss';

export class TwKeylabel extends HTMLElement {
	constructor() {
		super();
	}

	connectedCallback(): void {
		console.log("Keylabel connected");
	}
}

customElements.define("tw-keylabel", TwKeylabel);