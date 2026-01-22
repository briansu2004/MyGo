import { Component, signal } from '@angular/core';
import { CommonModule } from '@angular/common';

@Component({
  selector: 'app-root',
  standalone: true,
  imports: [CommonModule],
  template: `
    <div style="padding: 2rem; font-family: 'Segoe UI', sans-serif;">
      <h1 style="color: #007bff;">Go + Angular: Concurrency Demo</h1>

      <div style="margin-bottom: 20px;">
        <button
          (click)="startProcessing()"
          [disabled]="loading()"
          style="padding: 10px 25px; cursor: pointer; border-radius: 5px; border: none; background: #28a745; color: white;"
        >
          {{
            loading() ? 'Processing with Goroutines...' : 'Start Parallel Jobs'
          }}
        </button>
      </div>

      <div class="terminal">
        <div *ngFor="let log of logs()" class="log-line">
          <span class="timestamp">[{{ getNow() }}]</span> {{ log }}
        </div>
        <div *ngIf="logs().length === 0" style="color: #666;">
          Waiting for execution...
        </div>
      </div>
    </div>
  `,
  styles: [
    `
      .terminal {
        background: #1e1e1e;
        color: #d4d4d4;
        padding: 15px;
        border-radius: 8px;
        height: 300px;
        overflow-y: auto;
        font-family: monospace;
      }
      .log-line {
        margin-bottom: 5px;
        border-left: 3px solid #007bff;
        padding-left: 10px;
      }
      .timestamp {
        color: #858585;
        margin-right: 10px;
        font-size: 0.8rem;
      }
    `,
  ],
})
export class AppComponent {
  // Signals are the modern way to manage state in Angular
  logs = signal<string[]>([]);
  loading = signal<boolean>(false);

  startProcessing() {
    this.logs.set([]);
    this.loading.set(true);

    // Using EventSource for the Go SSE Stream
    const eventSource = new EventSource('http://localhost:8091/process');

    eventSource.onmessage = (event) => {
      // Update the signal with the new message from the Go worker
      this.logs.update((prev) => [...prev, event.data]);
    };

    eventSource.onerror = () => {
      eventSource.close();
      this.loading.set(false);
      this.logs.update((prev) => [
        ...prev,
        '✔️ BATCH COMPLETE: All Goroutines finished successfully.',
      ]);
    };
  }

  getNow() {
    return new Date().toLocaleTimeString();
  }
}
