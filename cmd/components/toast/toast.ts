class Toast extends HTMLElement {
	constructor() {
		super();
	}

	connectedCallback(): void {
		console.log("Toast connected");
	}
}

customElements.define("tw-toast", Toast);