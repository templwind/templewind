import './speed-dial.scss';

export class TwSpeedDial extends HTMLElement {
	constructor() {
		super();
	}

	connectedCallback(): void {
		console.log("SpeedDial connected");
	}
}

customElements.define("tw-speed-dial", TwSpeedDial);