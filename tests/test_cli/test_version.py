import unittest

from mock import patch

from click.testing import CliRunner

from polyaxon_cli.cli.version import upgrade, version


class TestVersion(unittest.TestCase):

    def setUp(self):
        self.runner = CliRunner()

    @patch('polyaxon_cli.cli.version.pip_upgrade')
    @patch('polyaxon_cli.cli.version.sys')
    def test_upgrade(self, mock_sys, pip_upgrade):
        mock_sys.version = '2.7.13 (default, Jan 19 2017, 14:48:08) \n[GCC 6.3.0 20170118]'
        self.runner.invoke(upgrade)
        pip_upgrade.assert_called_once()

    @patch('polyaxon_client.version.VersionClient.get_cli_version')
    @patch('polyaxon_cli.cli.version.dict_tabulate')
    def test_version_cli_default(self, dict_tabulate, get_cli_version):
        self.runner.invoke(version)
        get_cli_version.assert_called_once()
        dict_tabulate.assert_called_once()

    @patch('polyaxon_client.version.VersionClient.get_cli_version')
    @patch('polyaxon_cli.cli.version.dict_tabulate')
    def test_version_cli(self, dict_tabulate, get_cli_version):
        self.runner.invoke(version, ['--cli'])
        get_cli_version.assert_called_once()
        dict_tabulate.assert_called_once()

    @patch('polyaxon_client.version.VersionClient.get_platform_version')
    @patch('polyaxon_cli.cli.version.dict_tabulate')
    def test_version(self, dict_tabulate, get_version):
        self.runner.invoke(version, ['--platform'])
        get_version.assert_called_once()
        dict_tabulate.assert_called_once()
