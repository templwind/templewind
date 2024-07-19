import './radio.scss';

export class TwRadio extends HTMLElement {
	constructor() {
		super();
	}

	connectedCallback(): void {
		console.log("Radio connected");
	}
}

customElements.define("tw-radio", TwRadio);