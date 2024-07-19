import './datepicker.scss';

export class TwDatepicker extends HTMLElement {
	constructor() {
		super();
	}

	connectedCallback(): void {
		console.log("Datepicker connected");
	}
}

customElements.define("tw-datepicker", TwDatepicker);