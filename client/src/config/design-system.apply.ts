import { isPlatformBrowser } from '@angular/common';
import { IDesignSystem } from '../types/design-system.types';

export function applyDesignSystem(ds: IDesignSystem, platform_id: object): void {
  if (!isPlatformBrowser(platform_id)) return;
  const root = document.documentElement;

  root.style.setProperty('--subtitle-text-weight', ds.typography.subtitleTextWeight.toString());
  root.style.setProperty('--theme', ds.color.theme === 'system' ? 'light dark' : ds.color.theme);
}
