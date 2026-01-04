import { inject, Injectable } from '@angular/core';
import { TitleStrategy, RouterStateSnapshot } from '@angular/router';
import { Title } from '@angular/platform-browser';
import { DESIGN_SYSTEM } from '../config';

@Injectable()
export class AppTitleStrategy extends TitleStrategy {
  private readonly titleService = inject(Title);
  private readonly appName = inject(DESIGN_SYSTEM).app.APP_NAME;

  updateTitle(snapshot: RouterStateSnapshot): void {
    if (this.buildTitle(snapshot)) return;
    let pageTitle = snapshot.url.split('/').pop() || 'Home';
    pageTitle = pageTitle.charAt(0).toUpperCase() + pageTitle.slice(1);
    this.titleService.setTitle(`${this.appName} | ${pageTitle}`);
  }
}
