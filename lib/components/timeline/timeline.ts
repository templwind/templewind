import './timeline.scss';

export class TwTimeline extends HTMLElement {
	constructor() {
		super();
	}

	connectedCallback(): void {
		console.log("Timeline connected");
	}
}

customElements.define("tw-timeline", TwTimeline);