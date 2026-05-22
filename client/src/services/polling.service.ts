import { inject, Injectable } from '@angular/core';
import { DESIGN_SYSTEM } from '../config';

type PollKey = 'market' | 'trading';

@Injectable({ providedIn: 'root' })
export class PollingService {
  private ds = inject(DESIGN_SYSTEM);
  private timers = new Map<PollKey, ReturnType<typeof setInterval>>();
  private handlers = new Map<PollKey, () => void | Promise<void>>();

  register(key: PollKey, handler: () => void | Promise<void>): void {
    this.handlers.set(key, handler);
  }

  start(key: PollKey): void {
    if (this.timers.has(key)) return;
    const ms = key === 'trading' ? this.ds.devConfig.throttlingTimeMs : this.ds.devConfig.pollingTimeMs;
    const timer = setInterval(() => {
      void this.handlers.get(key)?.();
    }, ms);
    this.timers.set(key, timer);
  }

  stop(key: PollKey): void {
    const timer = this.timers.get(key);
    if (timer) {
      clearInterval(timer);
      this.timers.delete(key);
    }
  }

  stopAll(): void {
    for (const key of [...this.timers.keys()]) {
      this.stop(key);
    }
  }

  restart(key: PollKey): void {
    if (this.timers.has(key)) {
      this.stop(key);
      this.start(key);
    }
  }
}
