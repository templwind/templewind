import './tooltip.scss';

export class TwTooltip extends HTMLElement {
	constructor() {
		super();
	}

	connectedCallback(): void {
		console.log("Tooltip connected");
	}
}

customElements.define("tw-tooltip", TwTooltip);