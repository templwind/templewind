class DeviceMockups extends HTMLElement {
	constructor() {
		super();
	}

	connectedCallback(): void {
		console.log("DeviceMockups connected");
	}
}

customElements.define("tw-device-mockups", DeviceMockups);