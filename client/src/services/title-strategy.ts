import { inject, Injectable } from '@angular/core';
import { TitleStrategy, RouterStateSnapshot } from '@angular/router';
import { Title } from '@angular/platform-browser';

@Injectable()
export class AppTitleStrategy extends TitleStrategy {
  private readonly titleService = inject(Title);

  updateTitle(snapshot: RouterStateSnapshot): void {
    if (this.buildTitle(snapshot)) return;
    let pageTitle = snapshot.url.split('/').pop() || 'Home';
    pageTitle = pageTitle.charAt(0).toUpperCase() + pageTitle.slice(1);
    this.titleService.setTitle(`G-Market | ${pageTitle}`);
  }
}
