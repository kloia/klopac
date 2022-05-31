import unittest
from unittest.mock import Mock, MagicMock
from provisioner.repo import Repo


class TestInvalidRepo(unittest.TestCase):
    def setUp(self) -> None:
        data = {"uri": "", "state": {"enabled": ""}, "from_layer": ""}
        self.mock = MagicMock()
        self.mock.__getitem__.side_effect = data.__getitem__

    def test_empty_fields(self):
        with self.assertRaises(KeyError):
            Repo({}, "empty_repo")

    def test_empty_values(self):
        mock_repo = Repo(self.mock, "empty_uri")
        self.assertEqual(mock_repo.uri, "")
        self.assertEqual(mock_repo.enabled, "")
        self.assertEqual(mock_repo.layer, "")


if __name__ == "__main__":
    unittest.main()
