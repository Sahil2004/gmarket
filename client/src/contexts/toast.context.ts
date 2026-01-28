import { HttpContextToken } from '@angular/common/http';

export const SKIP_TOAST = new HttpContextToken<boolean>(() => false);
