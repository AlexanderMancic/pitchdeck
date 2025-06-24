class MyCounter extends HTMLElement {
  constructor() {
    super();
    this.count = 0;
    
    // Create shadow DOM
    this.attachShadow({ mode: 'open' });
    
    // Component HTML
    this.shadowRoot.innerHTML = `
      <style>
        :host {
          all: initial;
        }
        button {
          padding: 8px 16px;
          background: #3498db;
          color: white;
          border: none;
          border-radius: 4px;
          cursor: pointer;
          font-size: 1rem;
        }
        span {
          margin: 0 10px;
          font-weight: bold;
        }
      </style>
      <button id="decrement">-</button>
      <span id="count">0</span>
      <button id="increment">+</button>
    `;
    
    // Get elements
    this.incrementBtn = this.shadowRoot.getElementById('increment');
    this.decrementBtn = this.shadowRoot.getElementById('decrement');
    this.countDisplay = this.shadowRoot.getElementById('count');
    
    // Add event listeners
    this.incrementBtn.addEventListener('click', this.increment.bind(this));
    this.decrementBtn.addEventListener('click', this.decrement.bind(this));
  }
  
  increment() {
    this.count++;
    this.updateDisplay();
  }
  
  decrement() {
    this.count--;
    this.updateDisplay();
  }
  
  updateDisplay() {
    this.countDisplay.textContent = this.count;
    this.dispatchEvent(new CustomEvent('count-changed', { detail: this.count }));
  }
}

// Register the component
customElements.define('my-counter', MyCounter);
