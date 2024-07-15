class SpeedDial extends HTMLElement {
	constructor() {
		super();
	}

	connectedCallback(): void {
		console.log("SpeedDial connected");
	}
}

customElements.define("tw-speed-dial", SpeedDial);