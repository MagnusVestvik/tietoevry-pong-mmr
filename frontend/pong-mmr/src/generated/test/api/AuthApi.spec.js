/*
 * tietoevry-pong-mmr
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * OpenAPI spec version: 0.1
 *
 * NOTE: This class is auto generated by the swagger code generator program.
 * https://github.com/swagger-api/swagger-codegen.git
 *
 * Swagger Codegen version: 3.0.54
 *
 * Do not edit the class manually.
 *
 */
(function(root, factory) {
  if (typeof define === 'function' && define.amd) {
    // AMD.
    define(['expect.js', '../../src/index'], factory);
  } else if (typeof module === 'object' && module.exports) {
    // CommonJS-like environments that support module.exports, like Node.
    factory(require('expect.js'), require('../../src/index'));
  } else {
    // Browser globals (root is window)
    factory(root.expect, root.TietoevryPongMmr);
  }
}(this, function(expect, TietoevryPongMmr) {
  'use strict';

  var instance;

  beforeEach(function() {
    instance = new TietoevryPongMmr.AuthApi();
  });

  describe('(package)', function() {
    describe('AuthApi', function() {
      describe('apiLoginPost', function() {
        it('should call apiLoginPost successfully', function(done) {
          // TODO: uncomment, update parameter values for apiLoginPost call
          /*

          instance.apiLoginPost(body, function(error, data, response) {
            if (error) {
              done(error);
              return;
            }

            done();
          });
          */
          // TODO: uncomment and complete method invocation above, then delete this line and the next:
          done();
        });
      });
    });
  });

}));