class Radio extends HTMLElement {
	constructor() {
		super();
	}

	connectedCallback(): void {
		console.log("Radio connected");
	}
}

customElements.define("tw-radio", Radio);