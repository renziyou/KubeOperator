import {Pipe, PipeTransform} from '@angular/core';
import {TranslateService} from '@ngx-translate/core';

@Pipe({
    name: 'zoneStatus'
})
export class ZoneStatusPipe implements PipeTransform {

    constructor(private translateService: TranslateService) {
    }

    transform(value: string, ...args: unknown[]): unknown {
        let result = '';
        if (value) {
            switch (value) {
                case 'READY':
                    result = '<img src="assets/images/done.svg" border-style="none" vertical-align="middle">'
                        + this.translateService.instant('APP_STATUS_RUNNING');
                    break;
                case 'INITIALIZING':
                    result = '' + this.translateService.instant('APP_STATUS_INITIALING');
                    break;
                case 'UPLOADIMAGERROR':
                    result = '' + this.translateService.instant('APP_UPLOAD_IMAG_ERROR');
                    break;
                default:
                    result = value;
                    break;
            }
        }
        return result;
    }
}

