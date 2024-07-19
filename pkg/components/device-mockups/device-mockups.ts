import './device-mockups.scss';

export class TwDeviceMockups extends HTMLElement {
	constructor() {
		super();
	}

	connectedCallback(): void {
		console.log("DeviceMockups connected");
	}
}

customElements.define("tw-device-mockups", TwDeviceMockups);